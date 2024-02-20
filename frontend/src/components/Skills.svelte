<script lang="ts">
    import type { Skill } from '../models';
    import ProgrammingLanguage from './ProgrammingLanguage.svelte';
    export let skills: Skill[];
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
