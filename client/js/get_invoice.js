const client = require('./client');

client.invoiceClient.createInvoice({ number: "2" }, (error, inv) => {
    if (!error) {
        console.log('successfully fetch List notes');
        console.log("Response: ", inv);
    } else {
        console.error(error);
    }
});