import ClassCard from "./Cards/ClassCard";
import EventCard from "./Cards/EventCard";
import {TimelineItem} from "@/lib/api";


export default function RouteCard({item}: { item: TimelineItem }) {
  switch (item.type) {
    case "class":
    case "remote":
      return <ClassCard item={item}/>;
    case "event":
      return <EventCard item={item}/>;
    default:
      return null
  }
}