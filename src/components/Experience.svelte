<script lang="ts">
    import type { Experience } from '../models';

    const experience: Experience[] = [
        {
            company: 'Noiseless Acoustics',
            title: 'Senior Fullstack Developer',
            from: 'October 2020',
            to: 'Present',
            summary: 'Fullstack development (React + Django + PostgreSQL + lots of AWS microservices), DevOps, and a lot more',
        },
        {
            company: 'Analyse2',
            title: 'Fullstack Developer',
            from: 'October 2018',
            to: 'October 2020',
            summary: 'Fullstack development (Angular + asp.NET Core + MS SQL Server), architecture, database design, multiple external projects, DevOps, data analytics',
        },
        {
            company: 'QuattroFolia',
            title: 'Junior Software Developer',
            from: 'March 2017',
            to: 'October 2018',
            summary: 'Test engineering for Web and Android, Frontend development with AngularJS',
        },
        {
            company: 'Suomen Merivoimat',
            title: 'Military Service',
            from: 'January 2015',
            to: 'June 2015',
            summary: 'Fulfilled my military service duty at Upinniemi',
        },
        {
            company: 'Arc Technology Oy',
            title: 'Software Tester',
            from: 'October 2014',
            to: 'December 2014',
            summary: 'Manual testing of HRM & HRD software',
        },
    ];
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
