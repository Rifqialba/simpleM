export async function apiFetch<T>(
  path: string,
  options?: RequestInit,
): Promise<T> {

  const response = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL}${path}`,
    {
      ...options,
      headers: {
        "Content-Type": "application/json",
        ...(options?.headers || {}),
      },
    },
  )

  const data = await response.json()

  if (!response.ok) {
    throw new Error(
      data?.error?.message || "request failed",
    )
  }

  return data
}