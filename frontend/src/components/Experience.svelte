<script lang="ts">
    import type { Experience } from '../models';
    export let experience: Experience[];
</script>

<div>
    <h4>Experience</h4>
    <div class="experience">
        {#each experience as exp}
            <div class="col-3 col-sm-6">
                <h5>{exp.title}</h5>
                <div><b>{exp.company}</b> <small>({exp.from} - {exp.to})</small></div>
                <span>{exp.summary}</span>
            </div>
        {/each}
    </div>
</div>

<style lang="scss">
  .experience {
    display: flex;
    flex-wrap: wrap;

    b {
      font-weight: 700;
    }
  }
</style>
