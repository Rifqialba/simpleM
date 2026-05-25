"use client"

import { useState } from "react"

import { apiFetch } from "@/lib/api"

import { useAuthStore } from "@/stores/auth-store"

import { LoginResponse } from "@/types/auth"

export default function LoginPage() {

  const [email, setEmail] =
    useState("")

  const [password, setPassword] =
    useState("")

  const setToken =
    useAuthStore(
      (state) => state.setToken,
    )

  async function handleLogin() {

    try {

      const response =
        await apiFetch<LoginResponse>(
          "/login",
          {
            method: "POST",

            body: JSON.stringify({
              email,
              password,
            }),
          },
        )

      setToken(
        response.data.token,
      )

      alert("login success")

    } catch (error) {

      alert(
        error instanceof Error
          ? error.message
          : "login failed",
      )
    }
  }

  return (
    <main className="flex min-h-screen items-center justify-center">

      <div className="w-full max-w-sm space-y-4">

        <h1 className="text-2xl font-bold">
          Login
        </h1>

        <input
          className="w-full rounded border p-3"
          placeholder="Email"
          value={email}
          onChange={(e) =>
            setEmail(e.target.value)
          }
        />

        <input
          className="w-full rounded border p-3"
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) =>
            setPassword(e.target.value)
          }
        />

        <button
          onClick={handleLogin}
          className="w-full rounded bg-black p-3 text-white"
        >
          Login
        </button>

      </div>
    </main>
  )
}