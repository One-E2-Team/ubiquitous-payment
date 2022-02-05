export let WSserver = // #put "'" + WEBSHOP_SERVER_HOST_PORT + "'"
export let PSPserver = // #put "'" + PSP_SERVER_HOST_PORT + "'"
export let BankServer = // #put "'" + BANK_SERVER_HOST_PORT + "'"
  //export let WSserver = 'localhost:8001'
  //export let PSPserver = 'localhost:8002'
export let WSprotocol = // #put "'" + WEBSHOP_PROTOCOL + "'"
export let BankProtocol = // #put "'" + BANK_PROTOCOL + "'"

export function setJWTToken(jwt) {
  let new_roles = [];
  for (let item of jwt.roles) {
    new_roles.push(item.name);
  }
  jwt.roles = new_roles;
  sessionStorage.setItem("JWT", JSON.stringify(jwt));
}

export function getJWTToken() {
  return JSON.parse(sessionStorage.getItem("JWT"));
}

export function hasRole(role) {
  const token = JSON.parse(sessionStorage.getItem("JWT"));
  return token.roles.includes(role);
}

export function getHeader() {
  if (getJWTToken()) {
    return {
      Authorization: "Bearer " + getJWTToken().token
    };
  }
  return {
    Authorization: "Bearer "
  };
}

export function getUrlVars() {
  var vars = {};
  window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(m, key, value) {
    vars[key] = value;
  });
  return vars;
}
