import { twMerge as m } from "tailwind-merge";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    disabled?: boolean;
}

export default function Button({
    disabled = false,
    className,
    ...props
}: ButtonProps) {
    return (
        <button
            disabled={disabled}
            {...props}
            className={m(
                "flex h-12 items-center justify-center gap-1 rounded-xl px-6 py-1 transition-colors bg-black text-background font-semibold",
                disabled ? "opacity-50" : "",
                className,
            )}
        />
    );
}
