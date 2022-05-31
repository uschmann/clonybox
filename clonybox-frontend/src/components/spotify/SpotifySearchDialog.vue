<template>
  <v-dialog :value="isOpen" persistent max-width="800">
    <ToolbarCard title="Search spotify">
      <template #toolbar-items>
        <v-btn icon @click="close">
          <v-icon>
            mdi-close
          </v-icon>
        </v-btn>
      </template>

      <v-container fluid>
        <v-row>
          <v-col>
            <v-text-field label="Search"
                          :value="query"
                          @input="updateQuery($event)"
                          outlined dense/>
            <v-btn color="primary"
                   block
                   @click="search"
                   :loading="isLoading">
              search
            </v-btn>
          </v-col>
        </v-row>
      </v-container>

      <v-divider/>

      <v-card-actions>
        <v-btn text @click="close">
          Cancel
        </v-btn>
        <v-spacer/>
      </v-card-actions>
    </ToolbarCard>
  </v-dialog>
</template>

<script>
import {mapState, mapActions, mapMutations} from 'vuex'
import ToolbarCard from "@/components/utils/ToolbarCard";

export default {
  name: "SpotifySearchDialog",
  data() {
    return {
      isOpen: false
    }
  },
  components: {
    ToolbarCard
  },
  computed: {
      ...mapState('spotify/search', [
          'isLoading',
          'result',
          'query'
      ])
  },
  methods: {
    ...mapActions('spotify/search', [
      'search'
    ]),
    ...mapMutations('spotify/search', [
      'updateQuery'
    ]),
    open() {
      this.isOpen = true;
    },
    close() {
      this.isOpen = false;
    }
  }
}
</script>

<style scoped>

</style>