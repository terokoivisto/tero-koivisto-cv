<script lang="ts">
    import type { Skill } from '../models';
    import ProgrammingLanguage from './ProgrammingLanguage.svelte';
    export let skills: Skill[];
</script>

<div>
    <h4>Skills</h4>
    <div class="skill-list">
        {#each skills as skill}
            <ProgrammingLanguage skill={skill} />
        {/each}
    </div>
</div>

<style lang="scss">
    .skill-list {
      display: flex;
      flex-direction: row;
      gap: 1rem;
      justify-content: space-between;
      flex-wrap: wrap;

    }
</style>
