<script lang="ts">
    import CV from './pages/CV.svelte';
    import {onMount} from "svelte";
    import type {CVData} from "./models";

    let data: CVData;
    onMount(async () => {
        data = await (await fetch("http://localhost:7777/cv")).json() as CVData;
    })
</script>

<svelte:head>
    <title>{data?.name ?? '-'}</title>
</svelte:head>

<main>
    <CV data={data}/>
</main>

<style lang="scss">
  main {
    height: 100vh;
    width: 100vw;
    display: flex;
    justify-content: center;
    background: var(--gradient);
  }
</style>
