type ReadyBridgeMessage = {
  kind: "ready";
  sessionId: string;
};

const TRUSTED_PARENT_ORIGIN = "https://app.example.test";

export function sendReady(parentWindow: Window, sessionId: string): void {
  const message: ReadyBridgeMessage = { kind: "ready", sessionId };
  parentWindow.postMessage(message, TRUSTED_PARENT_ORIGIN);
}
