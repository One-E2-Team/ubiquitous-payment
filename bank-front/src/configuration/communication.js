export let WSserver = 'localhost:1080'
export let PSPserver = 'localhost:1081'
export let BankServer = 'localhost:10001'
  //export let WSserver = 'localhost:8001'
  //export let PSPserver = 'localhost:8002'
export let WSprotocol = 'http'
export let BankProtocol = 'http'

export function setJWTToken(jwt) {
  let new_roles = [];
  for (let item of jwt.roles) {
    new_roles.push(item.name);
  }
  jwt.roles = new_roles;
  sessionStorage.setItem("JWT", JSON.stringify(jwt));
}

export function logOut() {
  sessionStorage.removeItem("JWT");
}

export function getJWTToken() {
  if (sessionStorage.getItem("JWT") == null || sessionStorage.getItem("JWT") == undefined) {
    return null;
  }
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