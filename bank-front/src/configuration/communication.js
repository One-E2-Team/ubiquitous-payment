export let WSserver = 'localhost:1080'
export let PSPserver = 'localhost:1081'
export let BankServer = 'localhost:10001'
  //export let WSserver = 'localhost:8001'
  //export let PSPserver = 'localhost:8002'
export let WSprotocol = 'http'
export let BankProtocol = 'http'

export function getUrlVars() {
  var vars = {};
  window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(m, key, value) {
    vars[key] = value;
  });
  return vars;
}