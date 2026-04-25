// src/App.tsx
import { BrowserRouter, useRoutes } from "react-router-dom";
import routes from "~react-pages";
console.log("Generated routes:", routes);

function RoutesRenderer() {
    return useRoutes(routes);
}

function App() {
    return (
        <BrowserRouter>
            <RoutesRenderer />
        </BrowserRouter>
    );
}

export default App;
