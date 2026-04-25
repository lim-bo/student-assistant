import { Outlet } from "react-router-dom";
import Wrapper from "./wrapper";

export default function RootLayout() {
    return (
        <div className="min-h-screen bg-gray-50">
            <header>...</header>
            <main className="p-6">
                <Wrapper>
                    <Outlet />
                </Wrapper>
            </main>
        </div>
    );
}
