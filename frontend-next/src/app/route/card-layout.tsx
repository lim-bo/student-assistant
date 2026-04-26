import React from "react";

export function CardLayout({
    title,
    badge,
    children,
}: {
    title: React.ReactNode;
    badge?: React.ReactNode;
    children?: React.ReactNode;
}) {
    return (
        <li className="border rounded-xl p-3 bg-white shadow-sm">
            <div className="flex justify-between items-start gap-2">
                <h3 className="font-semibold basis-2/3">{title}</h3>

                {badge && (
                    <span className="text-sm text-gray-500 whitespace-nowrap shrink-0">
                        {badge}
                    </span>
                )}
            </div>

            {children && <div className="mt-2">{children}</div>}
        </li>
    );
}
