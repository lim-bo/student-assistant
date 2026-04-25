import RouteCard from "./route-card";
import { nanoid } from "nanoid";
import { TimelineItem } from "@/lib/api";

interface RouteListProps {
    children: TimelineItem[];
}

export default function RouteList({ children }: RouteListProps) {
    return (
        <ol className="space-y-2 mb-4">
            {children.map((item) => (
                <RouteCard key={nanoid()}>{item}</RouteCard>
            ))}
        </ol>
    );
}
