import { CardLayout } from "@/app/route/card-layout";
import { InPersonClass, RemoteClass, EventItem } from "@/lib/api";

interface EntryProps {
    children: InPersonClass | RemoteClass | EventItem;
}

export default function EntryCard({ children }: EntryProps) {
    return (
        <CardLayout
            title={children.title}
            badge={
                children.type === "remote" ? (
                    <span>онлайн</span>
                ) : children.type === "event" ? (
                    <span>событие</span>
                ) : (
                    <span>занятие</span>
                )
            }
        >
            <div className="text-sm text-gray-600">
                {children.time_start} – {children.time_end}
            </div>
            <div className="text-sm mt-2">
                {children.type === "remote" ? (
                    <span>онлайн</span>
                ) : children.type === "event" ? (
                    <>
                        <div className="text-sm text-gray-600">
                            {children.time_start} – {children.time_end}
                        </div>

                        <div className="text-sm mt-2">
                            📍 {children.location}
                        </div>
                    </>
                ) : (
                    <span>📍 {children.location}</span>
                )}
            </div>
        </CardLayout>
    );
}
