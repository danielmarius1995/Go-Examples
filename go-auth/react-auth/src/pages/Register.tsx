import React, { SyntheticEvent, useState } from "react";
import { Redirect } from "react-router-dom";

import endpoint from "../api/endpoint";

interface FormData {
  name: string;
  email: string;
  password: string;
}

const Register = () => {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);

  const submit = async (e: SyntheticEvent): Promise<void> => {
    e.preventDefault();

    // await fetch("http://localhost:4000/api/register", {
    //   method: "POST",
    //   headers: { "Content-Type": "application/json" },
    //   body: JSON.stringify({
    //     name,
    //     email,
    //     password
    //   })
    // });

    const body: FormData = {
      name,
      email,
      password
    };

    await endpoint.post("/register", { ...body });

    setRedirect(true);
  };

  if (redirect) {
    return <Redirect to="/login" />;
  }

  return (
    <form onSubmit={submit}>
      <h1 className="h3 mb-3 fw-normal">Please register</h1>

      <input
        className="form-control"
        placeholder="Name"
        required
        onChange={e => setName(e.target.value)}
      />

      <input
        type="email"
        className="form-control"
        placeholder="Email address"
        required
        onChange={e => setEmail(e.target.value)}
      />

      <input
        type="password"
        className="form-control"
        placeholder="Password"
        required
        onChange={e => setPassword(e.target.value)}
      />

      <button className="w-100 btn btn-lg btn-primary" type="submit">
        Register
      </button>
    </form>
  );
};

export default Register;
