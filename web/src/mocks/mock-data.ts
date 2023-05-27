const range = 10;
interface HappinessPoints {
    self_points: number;
    work_points: number;
    co_worker_points: number;
}

const randomNumber = (min: number = -5, max: number = 5): number => {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

const randomHappinessPoint = (): HappinessPoints => {
    return {
        self_points: randomNumber(),
        work_points: randomNumber(),
        co_worker_points: randomNumber(),
    }
}
interface HappinessHistory {
    happiness_points: HappinessPoints;
    date: string;
}
const generateData = (n: number): HappinessHistory[] => {
    const history: HappinessHistory[] = []
    for (let i = 0; i < n; i++) {
        const txn: HappinessHistory = {
            happiness_points: {
                ...randomHappinessPoint()
            },
            date: generateDate(i),
        }
        history.push(txn)
    }
    return history
}

const generateDate = (i: number): string => {
    const week = i * 7;
    const date = new Date(Date.now())
    const dateFormattedString = date.toLocaleDateString()
    return dateFormattedString
}

export const happinessPointsMock = {
    id: 1,
    period: "weekly",
    record: generateData(range)
}

const womanURL = "https://cdn-icons-png.flaticon.com/512/6997/6997662.png"
const menURL = "https://cdn-icons-png.flaticon.com/512/236/236831.png"
export const employees = [
    {
        id: 1,
        name: "Patchara Kleebbua",
        hp: 0,
        position: "Fullstack Developer",
        avatarURL: menURL
    },
    {
        id: 2,
        name: "Sayfar Hongsang",
        hp: 50,
        position: "Fullstack Developer",
        avatarURL: menURL
    },
    {
        id: 3,
        name: "Chonkamon Wongwaikunanan",
        hp: 0,
        position: "Project Manager",
        avatarURL: womanURL
    },
    {
        id: 4,
        name: "Pittaya Jindaruang",
        hp: 75,
        position: "Frontend Developer",
        avatarURL: menURL
    },
    {
        id: 5,
        name: "Kittipan Laokot",
        hp: 0,
        position: "Frontend Developer",
        avatarURL: menURL
    },
    {
        id: 6,
        name: "Bundit Korndee",
        hp: 100,
        position: "Project Owner",
        avatarURL: menURL
    },
    {
        id: 7,
        name: "Bike Chanokchon",
        hp: 90,
        position: "Fullstack Developer",
        avatarURL: menURL
    },
    {
        id: 8,
        name: "Jirat Dekum",
        hp: 20,
        position: "Fullstack Developer",
        avatarURL: menURL
    },
]

export const users = [
    {
        id: 1,
        email: "patchara.pck@gmail.com",
        name: "Patchara K.",
        password: "secret",
        position: "Human Resource"
    },
    {
        id: 2,
        email: "sayfar@gmail.com",
        name: "Sayfar H.",
        password: "secret",
        position: "Human Resource"
    },
]
// Profile URL
export const token = {
    accessToken: "acc123456789-abc",
    refreshToken: "ref-123456789abc",
}
// https://i.pinimg.com/originals/a6/58/32/a65832155622ac173337874f02b218fb.png Female
// https://cdn.icon-icons.com/icons2/2643/PNG/512/male_boy_person_people_avatar_icon_159358.png Male