const client = require('./client');
const startArg = process.argv[2];
var searchID;

// Soll gezielt einen Error hervorrufen
if (startArg === "error")
    searchID = 1;
else
    searchID = startArg;

client.productClient.getVolumeOfSales({ productId: searchID }, (error, res) => {
    if (!error) {
        console.log('success', res);

    } else {
        console.log('Error in getVolumeOfSales', error);
    }
});

