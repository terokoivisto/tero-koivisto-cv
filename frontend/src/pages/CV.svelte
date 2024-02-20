<script lang="ts">
    import { fade } from 'svelte/transition';
    import Skills from '../components/Skills.svelte';
    import Header from '../components/Header.svelte';
    import Experience from '../components/Experience.svelte';
    import type {CVData} from "../models";

    export let data: CVData;
</script>

{#if data}
    <div class="layout col-8" transition:fade>
        <Header name={data.name}
                title={data.title}
                aboutMe={data.aboutMe}
                personalMe={data.personalMe}
                location={data.location}
        />
        <Skills skills={data.skills}/>
        <Experience experience={data.experience}/>
        <small>This project was built using Svelte and Go. The source code is available at <a
                href="https://github.com/terokoivisto/tero-koivisto-cv">https://github.com/terokoivisto/tero-koivisto-cv</a></small>
    </div>
{/if}

<style lang="scss">
    .layout {
        display: flex;
        flex-direction: column;
        overflow: auto;
        gap: 2rem;

        background: var(--gradient-bg);
        border-radius: 1rem;
        margin: 1rem;
        padding: 1rem;
        filter: drop-shadow(0 0 0.66rem #3338);

        a {
            color: var(--primary);
        }
    }
</style>
