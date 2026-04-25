import EntryCard from "./entry-card";
import { TimelineItem } from "@/lib/api";

interface RouteCardProps {
    children: TimelineItem;
}

export default function RouteCard({ children }: RouteCardProps) {
    switch (children.type) {
        case "class":
        case "remote":
            return <EntryCard>{children}</EntryCard>;
        case "event":
            return <EntryCard>{children}</EntryCard>;
        default:
            return null;
    }
}
