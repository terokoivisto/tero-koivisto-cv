<script lang="ts">
    import type { Skill } from '../models';
    import ProgrammingLanguage from './ProgrammingLanguage.svelte';

    const skills: Skill[] = [
        {
            icon: 'vscode-icons:file-type-python',
            name: 'Python',
            usages: [
                'Django backend',
                'File processing',
                'AWS Lambdas',
                'Ad-hoc scripts',
                'Data analytics',
                'Much more...',
            ],
        },
        {
            icon: 'vscode-icons:file-type-typescript-official',
            name: 'Typescript',
            usages: [
                'React',
                'Svelte',
                'NodeJS',
                'AWS Lambdas',
                'FiveM',
                'Angular',
            ],
        },
        {
            icon: 'mdi:database',
            name: 'Databases',
            usages: [
                'PostgreSQL',
                'Designing database models/designs/architectures',
                'Dynamo DB',
                'MS SQL Server',
                'MariaDB',
                'MySQL',
            ],
        },
        {
            icon: 'skill-icons:aws-dark',
            name: 'AWS',
            usages: [
                "Lambda",
                "S3",
                "DynamoDB",
                "Aurora RDS",
                "Step functions",
                "AppSync",
                "ECR",
                "ECS",
                "Cloudformation",
                "CloudFront",
                "Load balancing",
                "Much more..."
            ],
        },
        {
            icon: "mingcute:hat-fill",
            name: "Other",
            usages: [
                "LUA",
                "Docker",
                "CI/CD",
                "Automated testing",
                "C#",
                "Java",
                "C/C++",
            ]
        }
    ];
</script>

<div>
    <h4>Skills</h4>
    <div class="language-list">
        {#each skills as skill}
            <div class="col-4 col-sm-6">
                <ProgrammingLanguage skill={skill} />
            </div>
        {/each}
    </div>
</div>

<style lang="scss">
</style>
