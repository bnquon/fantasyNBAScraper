"use client";
import { useRef } from "react";
import { Button, Input } from "@mui/material";

export default function Home() {
  const emailInput = useRef<HTMLInputElement>(null);

  const signup = async () => {
    const email = emailInput.current?.value;
    if (email) {
      try {
        const response = await fetch("/api/signup", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ email }),
        })

        if (response.ok) {
          alert("Email saved!");
        } else {
          alert("Error saving email");
        }
      } catch (error) {
        console.error(error);
      }
    }
  }

  return (
    <div className="flex justify-center items-center w-screen h-screen">
      <div className="bg-[#f5f5f5] w-[500px] h-[400px] flex flex-col p-8 items-center justify-center relative">
        <h1 className="text-2xl font-bold mb-4 text-center w-full absolute top-[10%]">
          Sign up to receive daily emails about the top fantasy performers!
        </h1>
        <Input
          inputRef={emailInput}
          type="email"
          placeholder="Email"
          sx={{ width: '80%', mt: 4, fontSize: '1.5rem' }}
        />
        <Button
          variant="contained"
          onClick={signup} // Trigger signup on button click
          sx={{ width: 'fit-content', mt: 4, position: 'absolute', bottom: '10%', fontSize: '1rem' }}
        >
          Sign up
        </Button>
      </div>
    </div>
  );
}
