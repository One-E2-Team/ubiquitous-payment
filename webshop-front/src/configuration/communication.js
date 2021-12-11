//export let WSserver = 'localhost:1080'
export let WSserver = 'localhost:8001'
export let WSprotocol = 'http'

export function setJWTToken(jwt) {
    let new_roles = [];
    for(let item of jwt.roles){
      new_roles.push(item.name);
    }
    jwt.roles = new_roles;
    sessionStorage.setItem("JWT", JSON.stringify(jwt));
}

export function getJWTToken() {
    return JSON.parse(sessionStorage.getItem("JWT"));
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