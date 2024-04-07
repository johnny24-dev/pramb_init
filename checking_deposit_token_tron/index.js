const axios = require('axios');

// Function to fetch all event logs for a contract within a block range
const fetchAllContractEventLogs = async (contractAddress, startBlock, endBlock) => {
    try {
        let allLogs = [];
        let page = 1;

        // Fetch logs for each page until all logs are fetched
        while (true) {
            const response = await axios.get(`https://apilist.tronscanapi.com/api/contract/event?contract=${contractAddress}&eventName=Transfer&fromTimestamp=${startBlock}&toTimestamp=${endBlock}&page=${page}`);
            const logs = response.data.data;
            if (logs.length === 0) break;
            allLogs = allLogs.concat(logs);
            page++;
        }

        return allLogs;
    } catch (error) {
        console.error(`Error fetching event logs for contract ${contractAddress} between blocks ${startBlock} and ${endBlock}:`, error);
        return [];
    }
};

// Function to process event logs for a contract
const processContractEventLogs = (events) => {
    events.forEach(event => {
        const sender = event.result.from;
        const receiver = event.result.to;
        const amount = event.result.value;

        console.log(`Token transfer from ${sender} to ${receiver} in contract ${event.contract_address}: ${amount}`);
    });
};

// Main function to track TRC20 token transfers
const trackTokenTransfers = async () => {
    try {
        const currentBlock = await axios.get('https://apilist.tronscan.org/api/block/latest')
        const endBlock = currentBlock.data.number;
        const startBlock = endBlock - 1000; // Start tracking from 1000 blocks before the current block

        // Example: USDT token contract address
        const usdtContractAddress = 'TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t';
        const events = await fetchAllContractEventLogs(usdtContractAddress, startBlock, endBlock);
        processContractEventLogs(events);

        // Example: Other well-known token contract addresses
        // const otherTokenContractAddresses = ['TOKEN_ADDRESS_1', 'TOKEN_ADDRESS_2', ...];
        // for (const tokenContractAddress of otherTokenContractAddresses) {
        //     const events = await fetchAllContractEventLogs(tokenContractAddress, startBlock, endBlock);
        //     processContractEventLogs(events);
        // }

        // TODO: You can expand your search by inspecting Transfer events for other tokens' contracts you come across
    } catch (error) {
        console.error('Error tracking TRC20 token transfers:', error);
    }
};

// Start tracking TRC20 token transfers
trackTokenTransfers();
