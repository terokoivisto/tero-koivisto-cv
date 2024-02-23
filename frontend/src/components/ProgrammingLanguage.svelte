<script lang="ts">
    import type { Skill } from '../models';

    export let skill: Skill;
</script>

<div class="language">
    <div class="heading">
        <iconify-icon icon={skill.icon} width="1.5rem"/>
        <h5>{skill.name}</h5>
    </div>
    <ul style="max-width: 250px;">
        {#each skill.usages as use}
            <li>{use}</li>
        {/each}
    </ul>
</div>

<style lang="scss">
  .heading {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 0.5rem;
  }

  li {
    border-left: 2px solid var(--primary);
    margin: 0.5rem 0.33rem;
    padding: 0 0.33rem;
  }
</style>
