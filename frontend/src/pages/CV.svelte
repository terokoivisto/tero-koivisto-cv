<script lang="ts">
    import {fade} from 'svelte/transition';
    import Skills from '../components/Skills.svelte';
    import Header from '../components/Header.svelte';
    import Experience from '../components/Experience.svelte';
    import type {CVData} from "../models";

    export let data: CVData;
</script>

{#if data}
    <div class="layout" transition:fade>
        <Header name={data.name}
                title={data.title}
                aboutMe={data.aboutMe}
                personalMe={data.personalMe}
                location={data.location}
        />
        <div class="content">
            <div class="child">
                <Experience experience={data.experience}/>
            </div>
            <div class="child">
                <Skills skills={data.skills}/>
            </div>
        </div>
        <small>This project was built using Svelte and Go. The source code is available at <a
                href="https://github.com/terokoivisto/tero-koivisto-cv">https://github.com/terokoivisto/tero-koivisto-cv</a></small>
    </div>
{/if}

<style lang="scss">
  .layout {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    overflow: auto;
    gap: 2rem;

    background: var(--gradient-bg);
    border-radius: 1rem;
    margin: 1rem;
    padding: 1rem;
    filter: drop-shadow(0 0 0.66rem #3338);

    .content {
      .child {
        padding: 1rem 0
      }
    }

    @media (min-width: 600px) {
      .content {
        display: flex;
        flex-direction: row;

        .child {
          width: 50%;
          padding: 1rem;
        }
      }
    }

    a {
      color: var(--primary);
    }
  }
</style>
