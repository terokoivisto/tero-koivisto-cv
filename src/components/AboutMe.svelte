<script lang="ts">
    export let name: string;
    export let title: string;
    export let aboutMe: string;
    export let personalMe: string;
</script>

<div>
    <h2>{name}</h2>
    <h3>{title}</h3>
    <br/>
    <h4>About ME</h4>
    <div class="about-me">
        <span>{aboutMe}</span>
        <span>{personalMe}</span>
    </div>
</div>

<style lang="scss">
    .about-me {
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }
</style>

