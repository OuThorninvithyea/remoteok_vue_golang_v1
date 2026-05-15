import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useSidebarTestimonials() {
    const testimonialsData = ref([])
    const founderData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/sidebar')
            testimonialsData.value = res.data.testimonials
            founderData.value = res.data.founder
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { testimonialsData, founderData, loading, error }
}
