<template>
  <div class="chart-tile">
    <svg id="line"></svg>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import * as d3 from 'd3';

// const props = defineProps(['data'])
const data = [
  { date: '24-Apr-07', amount: 93.24 },
  { date: '25-Apr-07', amount: 95.35 },
  { date: '26-Apr-07', amount: 98.84 },
  { date: '27-Apr-07', amount: 99.92 },
  { date: '30-Apr-07', amount: 99.8 },
  { date: '1-May-07', amount: 99.47 },
  { date: '2-May-07', amount: 100.39 },
  { date: '3-May-07', amount: 100.4 },
  { date: '4-May-07', amount: 100.81 },
  { date: '7-May-07', amount: 103.92 },
  { date: '8-May-07', amount: 105.06 },
  { date: '9-May-07', amount: 106.88 },
  { date: '10-May-07', amount: 107.34 },
]

onMounted(() => {
  const width = 400;
  const height = 250;
  const marginTop = 20;
  const marginRight = 20;
  const marginBottom = 30;
  const marginLeft = 40;

  const svg = d3.select('#line').attr('width', width).attr('height', height)
  const g = svg.append('g')
  const parseTime = d3.timeParse('%d-%b-%y')

  const x = d3
    .scaleTime()
    .domain(
      d3.extent(data, function (d) {
        return parseTime(d.date)
      })
    )
    .rangeRound([marginLeft, width - marginRight])
  const y = d3
    .scaleLinear()
    .domain(
      d3.extent(data, function (d) {
        return d.amount
      })
    )
    .rangeRound([height - marginBottom, marginTop])

  const line = d3
    .line()
    .x(function (d) {
      return x(parseTime(d.date))
    })
    .y(function (d) {
      return y(d.amount)
    })

  g.append('g')
    .attr('transform', `translate(0,${height - marginBottom})`)
    .call(d3.axisBottom(x))

  g.append('g').attr('transform', `translate(${marginLeft},0)`).call(d3.axisLeft(y))

  g.append('path')
    .datum(data)
    .attr('fill', 'none')
    .attr('stroke', 'steelblue')
    .attr('stroke-width', 1.5)
    .attr('d', line)
});
</script>

<style lang="scss" scoped>
.chart-tile {
  margin: 1em;
  padding: 1em;
  border: 1px solid black;
}
</style>
