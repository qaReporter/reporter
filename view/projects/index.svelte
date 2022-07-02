<script>
    export let prjs = []
  </script>
  
  <h1>Projects</h1>
  
  <ul>
      {#each prjs as project}
      <li>{project.name} - {project.description}</li>
      {/each}
  </ul>