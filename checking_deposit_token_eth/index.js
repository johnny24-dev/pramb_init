const { ethers } = require('ethers');
const fs = require('fs');
const cron = require('node-cron');
// Define the Ethereum node endpoint
const nodeURL = 'https://mainnet.infura.io/v3/c3372c21cf894fab8c75d70a864b6078';

// Connect to Ethereum node
const provider = new ethers.JsonRpcProvider(nodeURL)

// ERC20 Transfer event signature
const transferEventSignature = 'Transfer(address,address,uint256)';


// ERC20 contract ABI
const erc20ABI = [
    'event Transfer(address indexed from, address indexed to, uint256 value)'
];

// Function to check if the job is currently running
const isJobRunning = () => {
    return fs.existsSync('job.lock');
};


const fetch_transfer = async () => {

    // Check if the job is already running
    if (isJobRunning()) {
        console.log('Job is already running.');
        return;
    }

    // Create a lock file to indicate that the job is running
    fs.writeFileSync('job.lock', '');

    // Retrieve the last processed block number and transaction hash
    let lastBlockNumber = 0;
    let lastTxHash = '';

    try {
        const resumeData = JSON.parse(fs.readFileSync('resumeData.json', 'utf8'));
        lastBlockNumber = +resumeData.lastBlockNumber || 0;
        lastTxHash = resumeData.lastTxHash || '';
    } catch (error) {
        console.error('Error reading resume data:', error);
    }

    // Specify the start and end block numbers
    const startBlockNumber = lastBlockNumber;
    const endBlockNumber = await provider.getBlockNumber();
    for (let blockNumber = startBlockNumber; blockNumber <= endBlockNumber; blockNumber++) {
        try {

            const block = await provider.getBlock(blockNumber)
            // Iterate through transactions in the block
            console.log('total_transactions: ', block.transactions.length)
            let continute_at = lastTxHash ? block.transactions.findIndex((tx) => tx == lastTxHash) : -1
            let count = continute_at;
            for (let i = 0; i < block.transactions.length; i++) {

                let txHash = block.transactions[i]
                console.log("🚀 ~ constfetch_transfer= ~ txHash:", txHash)
                console.log("🚀 ~ constfetch_transfer= ~ continute_at:", continute_at)
                // checking for resume at txn in lastblock
                if (blockNumber == startBlockNumber && i <= continute_at) continue

                try {

                    // Fetch transaction details
                    const tx = await provider.getTransaction(txHash);
                    // console.log("🚀 ~ constfetch_transfer= ~ tx:", tx)

                    // check native token
                    // Log transaction details
                    console.log('Block:', blockNumber);
                    console.log('Transaction:');
                    console.log('Hash:', tx.hash);
                    console.log('From:', tx.from);
                    console.log('To:', tx.to);
                    console.log('Value:', ethers.formatEther(tx.value), 'ETH');


                    console.log('__________________erc20____________________________')

                    const receipt = await provider.getTransactionReceipt(tx.hash);

                    // Check if the transaction has logs (events)
                    if (receipt.logs.length > 0) {
                        // Iterate through logs to filter Transfer events
                        for (const log of receipt.logs) {
                            // Decode log data to check for Transfer event
                            const iface = new ethers.Interface(erc20ABI);
                            const parsedLog = iface.parseLog(log);
                            //   console.log("🚀 ~ constfetch_transfer= ~ parsedLog:", parsedLog)

                            // Check if the log matches the Transfer event signature
                            if (log.topics[0] === ethers.id(transferEventSignature)) {
                                // Log ERC20 transfer details
                                console.log('Block:', blockNumber);
                                console.log('Transaction:');
                                console.log('Hash:', tx.hash);
                                console.log('Erc20 address:', log.address)
                                console.log('From:', parsedLog.args[0]);
                                console.log('To:', parsedLog.args[1]);
                                console.log('Value:', parsedLog.args[2].toString());
                            }
                        }
                    }
                    console.log('__________________end____________________________')
                    console.log('count', count)
                    count += 1;
                } catch (error) {
                    console.error('Error processing transaction:', error);
                    // Continue to the next transaction
                    continue;
                }
                // Update the resume data with the latest block number and transaction hash
                const resumeData = JSON.stringify({ lastBlockNumber: blockNumber.toString(), lastTxHash: txHash });
                fs.writeFileSync('resumeData.json', resumeData);
            }

        } catch (error) {
            console.error('Error fetching block details:', error);
            // Continue to the next block
            continue;
        }
    }

 // Remove the lock file to indicate that the job has finished
 fs.unlinkSync('job.lock');

}
// Schedule the job to run every 5 minutes
cron.schedule('*/1 * * * *', async () => {
    console.log('Starting ERC20 transfer tracking job...');
    await fetch_transfer();
    console.log('ERC20 transfer tracking job completed.');
});