fetch("services").then(res => {
    return res.json();
}).then((service) => {
    if (service.proxy === "true") {
        document.getElementById('proxy').classList.remove('has-text-success');
        document.getElementById('proxy').classList.add('has-text-danger');
    }
    if (service.proxy_class === "true") {
        document.getElementById('proxy_class').classList.remove('has-text-success');
        document.getElementById('proxy_class').classList.add('has-text-danger');
    }
    if (service.proxy_kc === "true") {
        document.getElementById('proxy_kc').classList.remove('has-text-success');
        document.getElementById('proxy_kc').classList.add('has-text-danger');
    }
    if (service.proxy_dc === "true") {
        document.getElementById('proxy_dc').classList.remove('has-text-success');
        document.getElementById('proxy_dc').classList.add('has-text-danger');
    }
    if (service.proxy_sc === "true") {
        document.getElementById('proxy_sc').classList.remove('has-text-success');
        document.getElementById('proxy_sc').classList.add('has-text-danger');
    }
    if (service.mail === "true") {
        document.getElementById('mail').classList.remove('has-text-success');
        document.getElementById('mail').classList.add('has-text-danger');
    }
    if (service.mx === "true") {
        document.getElementById('mx').classList.remove('has-text-success');
        document.getElementById('mx').classList.add('has-text-danger');
    }
    if (service.web === "true") {
        document.getElementById('web').classList.remove('has-text-success');
        document.getElementById('web').classList.add('has-text-danger');
    }
    if (service.vpn === "true") {
        document.getElementById('vpn').classList.remove('has-text-success');
        document.getElementById('vpn').classList.add('has-text-danger');
    }
}).catch(err => {
    throw err;
});