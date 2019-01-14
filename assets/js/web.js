fetch("web").then(res => {
    return res.json();
}).then((data) => {
    if (data.status === true) {
        document.getElementById('web').classList.remove('has-text-success');
        document.getElementById('web').classList.add('has-text-danger');
    }
}).catch(err => {
    throw err;
});