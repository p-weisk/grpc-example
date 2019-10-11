const client = require('./client');
const startArg = process.argv[2];

const invoiceId = startArg==="" ? 1 : startArg;

client.invoiceClient.findInvoiceById({ number: "" + invoiceId }, (error, inv) => {
    if (!error) {
        console.log("success:", inv);
    } else {
        console.error(error);
    }
});