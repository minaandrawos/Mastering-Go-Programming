namespace go hydraThrift

struct CrewMember{
    1: i32 id,
    2: string name,
    3: i32 secClearance,
    4: string position,
}

struct Ship{
    1: string shipname,
    2: string CaptainName,
    3: list<CrewMember> Crew,
}

service HydraThriftService {
    void addShip(1: Ship s)
}