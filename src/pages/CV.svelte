<script lang="ts">
    import Skills from '../components/Skills.svelte';
    import Header from '../components/Header.svelte';
    import Experience from '../components/Experience.svelte';
</script>

<div class="layout col-8">
    <Header/>
    <Skills/>
    <Experience/>
    <small>This project was built using Svelte. The source code is available at <a href="https://github.com/terokoivisto/tero-koivisto-cv">https://github.com/terokoivisto/tero-koivisto-cv</a></small>
</div>

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
