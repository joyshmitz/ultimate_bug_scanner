async function refreshCache(): Promise<void> {
  const response = await fetch("/api/cache/refresh");
  if (!response.ok) {
    throw new Error("cache refresh failed");
  }
}

export function scheduleRefresh(): number {
  return window.setTimeout(async () => {
    await refreshCache();
  }, 30000);
}
