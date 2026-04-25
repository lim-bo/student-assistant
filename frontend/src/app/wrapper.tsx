import type { ReactNode } from "react";

export default function Wrapper({ children }: { children: ReactNode }) {
    return <div className="max-w-4xl mx-auto px-4">{children}</div>;
}
