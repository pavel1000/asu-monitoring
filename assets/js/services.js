fetch("services").then(res => {
    return res.json();
}).then((service) => {
    if (service.proxy === "false") {
        document.getElementById('proxy').classList.remove('has-text-success');
        document.getElementById('proxy').classList.add('has-text-danger');
    }
    if (service.proxy_class === "false") {
        document.getElementById('proxy_class').classList.remove('has-text-success');
        document.getElementById('proxy_class').classList.add('has-text-danger');
    }
    if (service.proxy_kc === "false") {
        document.getElementById('proxy_kc').classList.remove('has-text-success');
        document.getElementById('proxy_kc').classList.add('has-text-danger');
    }
    if (service.proxy_dc === "false") {
        document.getElementById('proxy_dc').classList.remove('has-text-success');
        document.getElementById('proxy_dc').classList.add('has-text-danger');
    }
    if (service.proxy_sc === "false") {
        document.getElementById('proxy_sc').classList.remove('has-text-success');
        document.getElementById('proxy_sc').classList.add('has-text-danger');
    }
    if (service.mail === "false") {
        document.getElementById('mail').classList.remove('has-text-success');
        document.getElementById('mail').classList.add('has-text-danger');
    }
    if (service.mx === "false") {
        document.getElementById('mx').classList.remove('has-text-success');
        document.getElementById('mx').classList.add('has-text-danger');
    }
    if (service.vpn === "false") {
        document.getElementById('vpn').classList.remove('has-text-success');
        document.getElementById('vpn').classList.add('has-text-danger');
    }
    var web = JSON.parse(service.web)
    if (web.status === 0) {
        document.getElementById('web').classList.remove('has-text-success');
        document.getElementById('web').classList.add('has-text-danger');
    }
    if (web.status === 1) {
        document.getElementById('web').classList.remove('has-text-success');
        document.getElementById('web').classList.add('has-text-warning');
    }
}).catch(err => {
    throw err;
});