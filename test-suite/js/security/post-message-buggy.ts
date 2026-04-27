type AuthBridgeMessage = {
  kind: "auth-token";
  token: string;
};

export function sendAuthToken(parentWindow: Window, token: string): void {
  const message: AuthBridgeMessage = { kind: "auth-token", token };
  parentWindow.postMessage(message, "*");
}
