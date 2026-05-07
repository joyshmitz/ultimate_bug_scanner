let leakedInterval;
let leakedObserver;

function parseUserInput(raw) {
  const parsedNumber = parseInt(raw);
  const parsedJson = JSON.parse(raw);

  if (parsedNumber === NaN) {
    throw "not a number";
  }

  try {
    doOptionalWork(parsedJson);
  } catch (err) {}

  new Function("payload", "return payload.admin === true")(parsedJson);
  Promise.resolve(parsedJson).then((value) => audit(value));
  callApiPromise(parsedJson);

  leakedInterval = setInterval(() => audit(parsedJson));
  leakedObserver = new MutationObserver(() => audit(parsedJson));

  return parsedJson;
}

function doOptionalWork(value) {
  return value && value.optional;
}

function audit(value) {
  return value;
}

function callApiPromise(value) {
  return Promise.resolve(value);
}
