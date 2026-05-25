"use client"

import { useEffect } from "react"

import {
  useParams,
} from "next/navigation"

import {
  useAuthStore,
} from "@/stores/auth-store"

import {
  useRealtimeStore,
} from "@/stores/realtime-store"

export default function RoomPage() {

  const params =
    useParams()

  const roomId =
    params.roomId as string

  const token =
    useAuthStore(
      (state) => state.token,
    )

  const connect =
    useRealtimeStore(
      (state) => state.connect,
    )

  const events =
    useRealtimeStore(
      (state) => state.events,
    )

  useEffect(() => {

    if (!token) {
      return
    }

    connect(
      roomId,
      token,
    )

  }, [
    roomId,
    token,
    connect,
  ])

  return (
    <main className="p-6">

      <h1 className="mb-4 text-2xl font-bold">
        Room: {roomId}
      </h1>

      <div className="space-y-2">

        {events.map(
          (event, index) => (

            <pre
              key={index}
              className="overflow-auto rounded bg-zinc-100 p-3 text-sm"
            >
              {JSON.stringify(
                event,
                null,
                2,
              )}
            </pre>
          ),
        )}

      </div>
    </main>
  )
}