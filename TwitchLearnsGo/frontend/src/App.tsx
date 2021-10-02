import React, { useState } from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import { TestForm } from "./components/FormTest";
import { RegisterForm } from "./components/RegisterForm";
import { Home } from "./components/Home";

function App() {
  const [count, setCount] = useState(0);

  return (
    <Router>
      <Switch>
        <Route exact path="/form" component={TestForm} />
        <Route exact path="/register" component={RegisterForm} />
        <Route exact path="/" component={Home} />
      </Switch>
    </Router>
  );
}

export default App;
