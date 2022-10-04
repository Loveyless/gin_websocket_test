export function useWebSocket(url: string, onMessage: (data: any) => void) {
  const ws = new WebSocket(url);

  ws.onopen = (e) => {
    console.log("open", e);
  };

  ws.onmessage = onMessage;

  ws.onclose = (e) => {
    console.log("close", e);
  };

  ws.onerror = (e) => {
    console.log("error", e);
  };

  return ws;
}
