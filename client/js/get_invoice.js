const client = require('./client');
client.findInvoiceById({number: "2"}, (error, inv) => {
    if (!error) {
        console.log('successfully fetch List notes');
        console.log(inv);
    } else {
        console.error(error);
    }
});