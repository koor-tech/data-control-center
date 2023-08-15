<template>
  <div class="chart-tile">
    <svg id="bar"></svg>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import * as d3 from 'd3'

// const props = defineProps(['data'])
const data = [
  { disc: 'vol 1', used: 24.2, capacity: 40, capacityUnits: 'GiB' },
  { disc: 'vol 2', used: 37.9, capacity: 40, capacityUnits: 'GiB' },
  { disc: 'vol 3', used: 45.7, capacity: 60, capacityUnits: 'GiB' },
  { disc: 'vol 4', used: 6.4, capacity: 60, capacityUnits: 'GiB' },
]

onMounted(() => {
  const width = 464
  const height = 250
  const marginTop = 30
  const marginRight = 0
  const marginBottom = 30
  const marginLeft = 40

  // Declare the x (horizontal position) scale.
  const x = d3
    .scaleBand()
    .domain(
      d3.groupSort(
        data,
        ([d]) => -d.disc,
        (d) => d.disc
      )
    ) // descending used
    .range([marginLeft, width - marginRight])
    .padding(0.1)

  // Declare the y (vertical position) scale.
  const y = d3
    .scaleLinear()
    .domain([0, 100])
    .range([height - marginBottom, marginTop])

  // Create the SVG container.
  const svg = d3
    .select('#bar')
    .attr('width', width)
    .attr('height', height)
    .attr('viewBox', [0, 0, width, height])
    .attr('style', 'max-width: 100%; height: auto;')

  // Add a rect for each bar.
  svg
    .append('g')
    .attr('fill', 'steelblue')
    .selectAll()
    .data(data)
    .join('rect')
    .attr('x', (d) => x(d.disc))
    .attr('y', (d) => y((d.used / d.capacity) * 100))
    .attr('height', (d) => y(0) - y((d.used / d.capacity) * 100))
    .attr('width', x.bandwidth())
    .attr('fill', (d) => ((d.used / d.capacity) * 100 > 80 ? 'firebrick' : null))

  // Add the x-axis and label.
  svg
    .append('g')
    .attr('transform', `translate(0,${height - marginBottom})`)
    .call(d3.axisBottom(x).tickSizeOuter(0))

  // Add the y-axis and label, and remove the domain line.
  svg
    .append('g')
    .attr('transform', `translate(${marginLeft},0)`)
    .call(d3.axisLeft(y))
    .call((g) => g.select('.domain').remove())
    .call((g) =>
      g
        .append('text')
        .attr('x', -marginLeft)
        .attr('y', 10)
        .attr('fill', 'currentColor')
        .attr('text-anchor', 'start')
        .text('↑ used (%)')
    )
})
</script>

<style lang="scss" scoped>
.chart-tile {
  margin: 1em;
  padding: 1em;
  border: 1px solid black;
}
</style>
