fetch("proxy").then(res => {
    return res.json();
}).then((data) => {
    if (data.status === true) {
        document.getElementById('proxy').classList.remove('has-text-success');
        document.getElementById('proxy').classList.add('has-text-danger');
    }
}).catch(err => {
    throw err;
});