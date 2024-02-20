export interface Skill {
    icon: string;
    name: string;
    usages: string[];
}

export interface Experience {
    company: string;
    title: string;
    from: string;
    to: string;
    summary: string;
}

export interface CVData {
    name: string;
    title: string;
    aboutMe: string;
    personalMe: string;
    location: string;
    skills: Skill[];
    experience: Experience[];
}
