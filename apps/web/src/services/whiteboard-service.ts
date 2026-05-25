import { apiFetch } from "@/lib/api"

import { WhiteboardState } from "@/types/whiteboard"

export async function getWhiteboardState(
  tabId: string,
  token: string,
) {

  return apiFetch<{
    data: WhiteboardState
  }>(
    `/tabs/${tabId}/whiteboard`,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    },
  )
}

export async function saveWhiteboardState(
  tabId: string,
  token: string,
  scene: any,
) {

  return apiFetch(
    `/tabs/${tabId}/whiteboard`,
    {
      method: "POST",

      headers: {
        Authorization: `Bearer ${token}`,
      },

      body: JSON.stringify({
        scene,
      }),
    },
  )
}