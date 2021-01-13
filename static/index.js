const WHEREIS_URI = "/api"
const RES_DOM_NODE = document.querySelector("#results");
const ERR_TEXT = "Error getting results from server..."
const WARN_TEXT = "Only http/https protocols supported. Prepedning https:// to url..."
const RES_LINK_CLASS = "res-link";
const USER_INPUT_SELECTOR = "#userInput";

function whereIsUri(userInput) {
  if (!(userInput.includes("http://") || 
    userInput.includes("https://"))) {
    const resDomNode = RES_DOM_NODE;
    const warnNode = document.createElement("h5");
    warnNode.textContent = WARN_TEXT;
    resDomNode.append(warnNode);
    userInput = `https://${userInput}`;
  }
  return `${WHEREIS_URI}?uri=${userInput}`
}

function fetchWhereIs(userInput) {
  return fetch(whereIsUri(userInput))
    .then(res => {
      if (res.status != 200) {
        throw new Error(ERR_TEXT);
      }
      return res.json();
    });
}

function displayResults(whereIsRes) {
  const resDomNode = RES_DOM_NODE;
  const origUriNode = document.createElement("a");
  const redirUriNode = document.createElement("a");
  origUriNode.innerText = `Original: ${whereIsRes.BaseUri}`;
  origUriNode.href = whereIsRes.BaseUri;
  origUriNode.className = RES_LINK_CLASS;
  redirUriNode.innerText = `Redirected to ${whereIsRes.RedirectedUri}`;
  redirUriNode.href = whereIsRes.RedirectedUri;
  redirUriNode.className = RES_LINK_CLASS;
  resDomNode.appendChild(origUriNode);
  resDomNode.appendChild(document.createElement("br"));
  resDomNode.appendChild(redirUriNode);
}

function displayError() {
  const resDomNode = RES_DOM_NODE;
  const errNode = document.createElement("h3");
  errNode.textContent = ERR_TEXT;
  resDomNode.append(errNode);
}

function clearResultsDiv() {
  const resDomNode = RES_DOM_NODE;
  while (resDomNode.firstChild) {
    resDomNode.removeChild(resDomNode.lastChild);
  }
}

function handleSubmit(event) {
  event.preventDefault();
  clearResultsDiv();
  const userInput = document.querySelector(USER_INPUT_SELECTOR).value;
  fetchWhereIs(userInput)
    .then(res => displayResults(res))
    .catch(err => displayError());
}

document.querySelector("#userForm").addEventListener("submit", handleSubmit);
document.querySelector("#clearButton").addEventListener("click", (e) => clearResultsDiv());