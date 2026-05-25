import { create } from "zustand"

import { RealtimeEvent } from "@/types/realtime"

type RealtimeStore = {
  socket: WebSocket | null

  events: RealtimeEvent[]

  connect: (
    roomId: string,
    token: string,
  ) => void

  sendEvent: (
    event: RealtimeEvent,
  ) => void
}

export const useRealtimeStore =
  create<RealtimeStore>((set, get) => ({

    socket: null,

    events: [],

    connect: (
      roomId,
      token,
    ) => {

      const socket =
        new WebSocket(
        `${process.env.NEXT_PUBLIC_WS_URL}/ws/${roomId}?token=${token}`,
    )

      socket.onopen = () => {

        socket.send(
          JSON.stringify({
            type: "ping",
          }),
        )
      }

      socket.onmessage = (
        event,
      ) => {

        const parsed =
          JSON.parse(
            event.data,
          )

        set((state) => ({
          events: [
            ...state.events,
            parsed,
          ],
        }))
      }

      socket.onerror = (
        error,
      ) => {

        console.error(
          "websocket error",
          error,
        )
      }

      set({
        socket,
      })
    },

    sendEvent: (
      event,
    ) => {

      const socket =
        get().socket

      if (
        socket &&
        socket.readyState ===
          WebSocket.OPEN
      ) {

        socket.send(
          JSON.stringify(event),
        )
      }
    },

  }))