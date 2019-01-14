fetch("mail").then(res => {
    return res.json();
}).then((data) => {
    if (data.status === true) {
        document.getElementById('mail').classList.remove('has-text-success');
        document.getElementById('mail').classList.add('has-text-danger');
    }
}).catch(err => {
    throw err;
});