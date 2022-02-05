//export let WSserver = 'localhost:8001'
//export let PSPserver = 'localhost:8002'
export let WSserver = 'localhost:1080'
export let PSPserver = 'localhost:1081'
export let Protocol = 'http'

export function setJWTToken(jwt) {
    let new_roles = [];
    for(let item of jwt.roles){
      new_roles.push(item.name);
    }
    jwt.roles = new_roles;
    sessionStorage.setItem("JWT", JSON.stringify(jwt));
}

export function logOut(){
  sessionStorage.removeItem("JWT");
}

export function getJWTToken() {
    if(sessionStorage.getItem("JWT") == null || sessionStorage.getItem("JWT") == undefined){
      return null;
    }
    return JSON.parse(sessionStorage.getItem("JWT"));
}

export function hasRole(role) {
  const token = JSON.parse(sessionStorage.getItem("JWT"));
  if (token == null || token == undefined){
    return false;
  }
  if(token.roles.includes(role)){
    return true;
  }
  return false;
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
