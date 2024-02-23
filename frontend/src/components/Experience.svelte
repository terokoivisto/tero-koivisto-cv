<script lang="ts">
    import type {Experience} from '../models';

    export let experience: Experience[];
</script>

<div>
    <h4>Experience</h4>
    <div class="experience-list">
        {#each experience as exp}
            <div class="experience">
                <h5 class="title">{exp.title}</h5>
                <div class="company">
                    <b>{exp.company}</b>
                </div>
                <small>({exp.from} - {exp.to})</small>
                <span>{exp.summary}</span>
            </div>
        {/each}
    </div>
</div>

<style lang="scss">
  .experience-list {
    display: flex;
    flex-wrap: wrap;
    gap: 2rem;

    .experience {
      display: flex;
      flex-direction: column;
      //max-width: 450px;

      b {
        font-weight: 700;
      }
    }
  }
</style>
