const client = require('./client');
const startArg = process.argv;
var invoice;
if (startArg.length < 5)
    console.log('Parameters missing! Please set clientID, ProductID and number');
else {
    invoice = {
        clientId: startArg[2],
        p: {
            productId: startArg[3]
        },
        number: +startArg[4]
    }

    client.invoiceClient.createInvoice(invoice, (error, inv) => {
        if (!error) {
            console.log('successfully created invoice');
            console.log("Response: ", inv);
        } else {
            console.error(error);
        }
    });
}


