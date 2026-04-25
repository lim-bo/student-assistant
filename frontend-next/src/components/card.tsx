import type { ReactNode } from "react";
import { twMerge as m } from "tailwind-merge";

interface CardProps extends React.HTMLAttributes<HTMLDivElement> {
    children: ReactNode;
    className?: string;
}

export default function Card({
    children = null,
    className,
    ...props
}: CardProps) {
    return (
        <div className={m("p-1 rounded-sm", className)} {...props}>
            {children}
        </div>
    );
}
